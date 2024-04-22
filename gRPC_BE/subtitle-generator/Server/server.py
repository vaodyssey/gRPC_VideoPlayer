import grpc
from concurrent import futures
import time
import binascii
import subGenProto_pb2_grpc as pb2_grpc
import subGenProto_pb2 as pb2
import logging
class SubtitleGeneratorService(pb2_grpc.SubtitleGeneratorServicer):
    def __init__(self, *args, **kwargs):
        pass
    def Generate(self,request,context):      
        try:
            videoBytes = request.video      
            result = {'video':bytes(videoBytes)}
            return pb2.OutputVideo(**result)
        except Exception as e:
            logging.error('Error at %s', 'division', exc_info=e)

def serve():
    choptions = [('grpc.max_send_message_length', 100*1024*1024),
                ('grpc.max_receive_message_length', 100*1024*1024)]
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10),options = choptions)
    pb2_grpc.add_SubtitleGeneratorServicer_to_server(SubtitleGeneratorService(), server)
    max_message_length = 10 * 1024 * 1024  # 10MB in bytes    
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()