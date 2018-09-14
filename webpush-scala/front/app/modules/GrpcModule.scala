package modules

import com.google.inject.AbstractModule
import io.grpc.ManagedChannelBuilder
import nokamoto.protobuf.PushSubscriptionServiceGrpc.PushSubscriptionServiceStub
import nokamoto.protobuf.PushSubscriptionServiceGrpc
import nokamoto.protobuf.WebpushServiceGrpc
import nokamoto.protobuf.WebpushServiceGrpc.WebpushServiceStub
import play.api.Configuration
import play.api.Environment

class GrpcModule(environment: Environment, configuration: Configuration)
    extends AbstractModule {
  private[this] def configureWebpushService(): Unit = {
    val host = configuration.get[String]("grpc.webpush.host")
    val port = configuration.get[Int]("grpc.webpush.port")
    val channel =
      ManagedChannelBuilder.forAddress(host, port).usePlaintext().build()
    bind(classOf[WebpushServiceStub])
      .toInstance(WebpushServiceGrpc.stub(channel))
  }

  private[this] def configurePushSubscriptionService(): Unit = {
    val host = configuration.get[String]("grpc.push-subscription.host")
    val port = configuration.get[Int]("grpc.push-subscription.port")
    val channel =
      ManagedChannelBuilder.forAddress(host, port).usePlaintext().build()
    bind(classOf[PushSubscriptionServiceStub])
      .toInstance(PushSubscriptionServiceGrpc.stub(channel))
  }

  override def configure(): Unit = {
    configureWebpushService()
    configurePushSubscriptionService()
  }
}
