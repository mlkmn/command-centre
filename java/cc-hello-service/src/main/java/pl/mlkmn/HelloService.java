package pl.mlkmn;

import io.grpc.stub.StreamObserver;
import net.devh.springboot.autoconfigure.grpc.server.GrpcService;

@GrpcService(HelloGrpc.class)
public class HelloService extends HelloGrpc.HelloImplBase {

    @Override
    public void greet(HelloOuterClass.HelloRequest request, StreamObserver<HelloOuterClass.HelloResponse>
            responseObserver) {
        HelloOuterClass.HelloResponse reply = HelloOuterClass.HelloResponse.newBuilder().setMessage("Hello " +
                "=============> " + request.getName()).build();
        responseObserver.onNext(reply);
        responseObserver.onCompleted();
    }
}
