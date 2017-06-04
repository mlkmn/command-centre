package pl.cc;

import io.grpc.Channel;
import net.devh.springboot.autoconfigure.grpc.client.GrpcClient;
import org.springframework.stereotype.Service;

@Service
public class UIHelloService {

    @GrpcClient("cc-hello-server")
    private Channel serverChannel;

    String sendMessage(String name) {
        HelloGrpc.HelloBlockingStub stub = HelloGrpc.newBlockingStub(serverChannel);
        HelloOuterClass.HelloRequest request = HelloOuterClass.HelloRequest.newBuilder().setName(name)
                .build();
        HelloOuterClass.HelloResponse response = stub.greet(request);
        return response.getMessage();
    }
}
