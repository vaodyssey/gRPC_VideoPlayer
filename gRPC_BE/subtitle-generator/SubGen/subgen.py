import time
import math
import ffmpeg
import os
from faster_whisper import WhisperModel

tempPath = f'{os.getcwd()}/SubGen'
subgenPath = tempPath.replace(os.sep, '/')

input_video = f"{subgenPath}/input.mp4"
input_video_name = input_video.replace(".mp4", "")
def extract_audio():
    extracted_audio = f"{subgenPath}/extracted-audio.wav"
    stream = ffmpeg.input(input_video)
    stream = ffmpeg.output(stream, extracted_audio)
    ffmpeg.run(stream, overwrite_output=True)
    return extracted_audio

def transcribe(audio):
    model = WhisperModel("small")
    segments, info = model.transcribe(audio)
    language = info[0]
    print("Transcription language", info[0])
    segments = list(segments)
    for segment in segments:
        # print(segment)
        print("[%.2fs -> %.2fs] %s" %
              (segment.start, segment.end, segment.text))
    return language, segments
def format_time(seconds):

    hours = math.floor(seconds / 3600)
    seconds %= 3600
    minutes = math.floor(seconds / 60)
    seconds %= 60
    milliseconds = round((seconds - math.floor(seconds)) * 1000)
    seconds = math.floor(seconds)
    formatted_time = f"{hours:02d}:{minutes:02d}:{seconds:01d},{milliseconds:03d}"

    return formatted_time
def generate_subtitle_file(language, segments):
    subtitle_file = f"{subgenPath}/subtitle.{language}.srt"
    text = ""
    for index, segment in enumerate(segments):
        segment_start = format_time(segment.start)
        segment_end = format_time(segment.end)
        text += f"{str(index+1)} \n"
        text += f"{segment_start} --> {segment_end} \n"
        text += f"{segment.text} \n"
        text += "\n"
        
    f = open(subtitle_file, "w")
    f.write(text)
    f.close()

    return subtitle_file
def add_subtitle_to_video(soft_subtitle, subtitle_file:str,  subtitle_language):

    video_input_stream = ffmpeg.input(input_video)
    subtitle_input_stream = ffmpeg.input(subtitle_file)
    output_video = f"{subgenPath}/output.mp4"
    subtitle_track_title = subtitle_file.replace(".srt", "")

    if soft_subtitle:
        stream = ffmpeg.output(
            video_input_stream, subtitle_input_stream, output_video, **{"c": "copy", "c:s": "mov_text"},
            **{"metadata:s:s:0": f"language={subtitle_language}",
            "metadata:s:s:0": f"title={subtitle_track_title}"}
        )
        ffmpeg.run(stream, overwrite_output=True)
    else:
        parsedSubFilePath = parseSubtitleFilePath(subtitle_file)
        stream = ffmpeg.output(video_input_stream, output_video,

                               vf=f"subtitles='{parsedSubFilePath}'")

        ffmpeg.run(stream, overwrite_output=True)
def parseSubtitleFilePath(path):
    drive, rest_of_path = os.path.splitdrive(path)        
    replaced_drive = drive.replace(':', '\:')            
    replaced_path = replaced_drive + rest_of_path    
    return replaced_path

def run():
    extracted_audio = extract_audio()
    language, segments = transcribe(audio=extracted_audio)
    subtitle_file = generate_subtitle_file(language=language,segments=segments)
    add_subtitle_to_video(
        soft_subtitle=False,
        subtitle_file=subtitle_file,
        subtitle_language=language       
    )

