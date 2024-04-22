
import sys
import os
from concurrent import futures
#Allows Python to see the Server folder containing server.py and gRPC modules .
subGenPath = os.getcwd()
sys.path.append(f'{subGenPath}/Server') 

import grpc
import Server.subGenProto_pb2_grpc as pb2_grpc
import Server.subGenProto_pb2 as pb2
from Server.server import SubtitleGeneratorService


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