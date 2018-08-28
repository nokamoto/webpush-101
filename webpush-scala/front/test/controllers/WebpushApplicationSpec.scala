package controllers

import org.scalatestplus.play.PlaySpec
import org.scalatestplus.play.guice.GuiceOneAppPerTest
import play.api.mvc.Results
import play.api.test.FakeRequest
import play.api.test.Helpers._

class WebpushApplicationSpec
    extends PlaySpec
    with GuiceOneAppPerTest
    with Results {
  "WebpushApplication#index" should {
    "return OK" in {
      val sut = app.injector.instanceOf[WebpushApplication]
      val actual = sut.index(FakeRequest())

      assert(status(actual) === OK)
    }
  }

  "WebpushApplication#subscribe" should {
    "be not implemented yet" in {
      val sut = app.injector.instanceOf[WebpushApplication]
      val actual = sut.subscribe(FakeRequest())

      assert(status(actual) === NOT_IMPLEMENTED)
    }
  }

  "WebpushApplication#unsubscribe" should {
    "be not implemented yet" in {
      val sut = app.injector.instanceOf[WebpushApplication]
      val actual = sut.unsubscribe(FakeRequest())

      assert(status(actual) === NOT_IMPLEMENTED)
    }
  }

  "WebpushApplication#testSubscription" should {
    "be not implemented yet" in {
      val sut = app.injector.instanceOf[WebpushApplication]
      val actual = sut.testSubscription(FakeRequest())

      assert(status(actual) === NOT_IMPLEMENTED)
    }
  }
}
