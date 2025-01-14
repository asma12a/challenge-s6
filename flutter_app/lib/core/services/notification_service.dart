import 'package:firebase_messaging/firebase_messaging.dart';
import 'package:flutter/cupertino.dart';


class NotificationService {

  final _firebaseMessaging = FirebaseMessaging.instance;

  Future<void> initNotifications() async {
    await _firebaseMessaging.requestPermission();

    final fcmToken = await _firebaseMessaging.getToken();

    debugPrint("Token $fcmToken");
  }



}