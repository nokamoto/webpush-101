package controllers

import com.google.protobuf.empty.Empty
import models.PushSubscription
import nokamoto.protobuf.PushSubscriptionNotification
import nokamoto.protobuf.User
import nokamoto.protobuf.UserSubscription
import nokamoto.protobuf.PushSubscriptionServiceGrpc.PushSubscriptionServiceStub
import nokamoto.protobuf.WebpushServiceGrpc.WebpushServiceStub
import org.mockito.ArgumentCaptor
import org.scalatest.mockito.MockitoSugar
import org.scalatestplus.play.PlaySpec
import play.api.mvc.Results
import play.api.test.FakeRequest
import play.api.test.Helpers._
import org.mockito.Mockito._
import org.mockito.Matchers._

import scala.concurrent.Future

class WebpushApplicationSpec extends PlaySpec with MockitoSugar with Results {
  private[this] val endpoint =
    "https://updates.push.services.mozilla.com:443/wpush/v2/gAAAAABbiL"
  private[this] val p256dh =
    "BFn48GVRHcj1rWMiuQ0E9LqHB5MWq0OEq8T5hhiO8Mjr3O+mv3BrNLtx+bhwaV8WRJ6kQg1Q0OrtTzIyduDSPQI="
  private[this] val auth = "xPae4nGweJlQkdLMUppMKQ=="
  private[this] val subscription =
    PushSubscription(endpoint = endpoint, p256dh = p256dh, auth = auth)

  "WebpushApplication#index" should {
    "return OK" in {
      val sut = new WebpushApplication(mock[WebpushServiceStub],
                                       mock[PushSubscriptionServiceStub],
                                       stubControllerComponents())
      val actual = sut.index()(FakeRequest())

      assert(status(actual) === OK)
    }
  }

  "WebpushApplication#subscribe" should {
    "subscribe and return OK" in {
      val stub = mock[PushSubscriptionServiceStub]
      when(stub.subscribe(any())).thenReturn(Future.successful(Empty()))

      val sut = new WebpushApplication(mock[WebpushServiceStub],
                                       stub,
                                       stubControllerComponents())
      val actual = sut.subscribe()(FakeRequest().withBody(subscription))

      assert(status(actual) === OK)

      val captor = ArgumentCaptor.forClass(classOf[UserSubscription])
      verify(stub, times(1)).subscribe(captor.capture())

      val passed = captor.getValue.subscription
      assert(passed.length === 1)
      assert(subscription.proto === passed.head)

      assert(User() === captor.getValue.getUser, "to be fixed")
    }
  }

  "WebpushApplication#unsubscribe" should {
    "unsubscribe and return OK" in {
      val stub = mock[PushSubscriptionServiceStub]
      when(stub.unsubscribe(any())).thenReturn(Future.successful(Empty()))

      val sut = new WebpushApplication(mock[WebpushServiceStub],
                                       stub,
                                       stubControllerComponents())
      val actual = sut.unsubscribe()(FakeRequest().withBody(subscription))

      assert(status(actual) === OK)

      val captor =
        ArgumentCaptor.forClass(classOf[nokamoto.protobuf.PushSubscription])
      verify(stub, times(1)).unsubscribe(captor.capture())

      assert(subscription.proto === captor.getValue)
    }
  }

  "WebpushApplication#testSubscription" should {
    "send a push subscription notification and return OK" in {
      val stub = mock[WebpushServiceStub]
      when(stub.sendPushSubscriptionNotification(any()))
        .thenReturn(Future.successful(Empty()))

      val sut = new WebpushApplication(stub,
                                       mock[PushSubscriptionServiceStub],
                                       stubControllerComponents())
      val actual = sut.testSubscription()(FakeRequest().withBody(subscription))

      assert(status(actual) === OK)

      val captor =
        ArgumentCaptor.forClass(classOf[PushSubscriptionNotification])
      verify(stub, times(1)).sendPushSubscriptionNotification(captor.capture())

      val passed = captor.getValue.subscription
      assert(passed.length === 1)
      assert(subscription.proto === passed.head)
    }
  }
}
