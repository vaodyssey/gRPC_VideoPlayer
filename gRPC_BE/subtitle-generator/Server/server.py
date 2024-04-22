
import subGenProto_pb2_grpc as pb2_grpc
import subGenProto_pb2 as pb2
import logging
from SubGen import subgen
import os

serverPath = os.getcwd()
class SubtitleGeneratorService(pb2_grpc.SubtitleGeneratorServicer):
    def __init__(self, *args, **kwargs):
        pass
    def Generate(self,request,context):      
        try:
            videoBytes = request.video 
            self.SaveVideoToSubGen(videoBytes)                 
            subgen.run()
            result = self.GetResultVideo()
            return pb2.OutputVideo(**result)
        except Exception as e:
            logging.error('Error at %s', 'division', exc_info=e)
    
    def SaveVideoToSubGen(self,videoBytes):
        os.chdir("..")
        with open(f"{serverPath}/SubGen/input.mp4",'wb') as inputVideo:
            inputVideo.write(videoBytes)
            inputVideo.close()
        os.chdir(f"{serverPath}/Server")
    
    def GetResultVideo(self):
        os.chdir("..")
        with open(f"{serverPath}/SubGen/output.mp4", 'rb') as f:
            mp4_bytes = f.read()
        os.chdir(f"{serverPath}/Server")
        return {'video':bytes(mp4_bytes)}


