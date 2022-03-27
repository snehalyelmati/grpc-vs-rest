from concurrent import futures
import logging
import grpc
import hello_pb2
import hello_pb2_grpc

port = 50052

class Greeter(hello_pb2_grpc.HelloServicer):
    def SayHello(self, request, context):
        # print("Hello!")
        return hello_pb2.HelloResponse(message="Hello, %s!" % request.name)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    hello_pb2_grpc.add_HelloServicer_to_server(Greeter(), server)
    server.add_insecure_port('[::]:'+str(port))
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    logging.basicConfig()
    print("Server listening at: " + str(port))
    serve()
