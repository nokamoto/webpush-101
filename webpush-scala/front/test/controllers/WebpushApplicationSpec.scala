package controllers

import com.google.protobuf.empty.Empty
import models.PushSubscription
import nokamoto.protobuf.PushSubscriptionNotification
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

  "WebpushApplication#index" should {
    "return OK" in {
      val sut = new WebpushApplication(mock[WebpushServiceStub],
                                       stubControllerComponents())
      val actual = sut.index()(FakeRequest())

      assert(status(actual) === OK)
    }
  }

  "WebpushApplication#subscribe" should {
    "be not implemented yet" in {
      val sut = new WebpushApplication(mock[WebpushServiceStub],
                                       stubControllerComponents())
      val actual = sut.subscribe()(FakeRequest())

      assert(status(actual) === NOT_IMPLEMENTED)
    }
  }

  "WebpushApplication#unsubscribe" should {
    "be not implemented yet" in {
      val sut = new WebpushApplication(mock[WebpushServiceStub],
                                       stubControllerComponents())
      val actual = sut.unsubscribe()(FakeRequest())

      assert(status(actual) === NOT_IMPLEMENTED)
    }
  }

  "WebpushApplication#testSubscription" should {
    "send a push subscription notification and return OK" in {
      val endpoint =
        "https://updates.push.services.mozilla.com:443/wpush/v2/gAAAAABbiL"
      val p256dh =
        "BFn48GVRHcj1rWMiuQ0E9LqHB5MWq0OEq8T5hhiO8Mjr3O+mv3BrNLtx+bhwaV8WRJ6kQg1Q0OrtTzIyduDSPQI="
      val auth = "xPae4nGweJlQkdLMUppMKQ=="
      val req =
        PushSubscription(endpoint = endpoint, p256dh = p256dh, auth = auth)

      val stub = mock[WebpushServiceStub]
      when(stub.sendPushSubscriptionNotification(any()))
        .thenReturn(Future.successful(Empty()))

      val sut = new WebpushApplication(stub, stubControllerComponents())
      val actual = sut.testSubscription()(FakeRequest().withBody(req))

      assert(status(actual) === OK)

      val captor =
        ArgumentCaptor.forClass(classOf[PushSubscriptionNotification])
      verify(stub, times(1)).sendPushSubscriptionNotification(captor.capture())

      val subscriptions = captor.getValue.subscription
      assert(subscriptions.length === 1)
      assert(req.proto === subscriptions.head)
    }
  }
}
