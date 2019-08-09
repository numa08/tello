import UIKit
import Flutter
import Tellogo

@UIApplicationMain
@objc class AppDelegate: FlutterAppDelegate {
  override func application(
    _ application: UIApplication,
    didFinishLaunchingWithOptions launchOptions: [UIApplicationLaunchOptionsKey: Any]?
  ) -> Bool {
    GeneratedPluginRegistrant.register(with: self)
    print("load tellogo version: \(TellogoVersion())")
    return super.application(application, didFinishLaunchingWithOptions: launchOptions)
  }
}
