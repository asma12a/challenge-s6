import 'package:dio/dio.dart';
import 'package:firebase_messaging/firebase_messaging.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_local_notifications/flutter_local_notifications.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/core/utils/constants.dart';
import 'package:squad_go/main.dart';

class NotificationService {
  final FirebaseMessaging _firebaseMessaging = FirebaseMessaging.instance;
  final FlutterLocalNotificationsPlugin _localNotificationsPlugin = FlutterLocalNotificationsPlugin();

  Future<void> initNotifications() async {
    await _firebaseMessaging.requestPermission();

    const AndroidInitializationSettings initializationSettingsAndroid = AndroidInitializationSettings('@drawable/background');
    const InitializationSettings initializationSettings = InitializationSettings(
      android: initializationSettingsAndroid,
    );
    await _localNotificationsPlugin.initialize(initializationSettings);

    final fcmToken = await _firebaseMessaging.getToken();
    debugPrint("FCM Token: $fcmToken");
    if (fcmToken != null){
      await _storeFcmToken(fcmToken);
    }

    FirebaseMessaging.onMessage.listen((RemoteMessage message) {
      debugPrint("Notification reçue en premier plan : ${message.notification?.title}");
      _showNotification(message);
    });

    FirebaseMessaging.onMessageOpenedApp.listen((RemoteMessage message) {
      debugPrint("Notification ouverte depuis l'arrière-plan : ${message.notification?.title}");
    });

    FirebaseMessaging.onBackgroundMessage(_firebaseMessagingBackgroundHandler);
  }

  static Future<void> _firebaseMessagingBackgroundHandler(RemoteMessage message) async {
    debugPrint("Notification reçue en arrière-plan : ${message.notification?.title}");
  }

  Future<void> _showNotification(RemoteMessage message) async {
    const AndroidNotificationDetails androidDetails = AndroidNotificationDetails(
      'default_channel',
      'Notifications',
      channelDescription: 'Ce canal est utilisé pour les notifications générales',
      importance: Importance.max,
      priority: Priority.high,
    );

    const NotificationDetails notificationDetails = NotificationDetails(
      android: androidDetails,
    );

    await _localNotificationsPlugin.show(
      0,
      message.notification?.title ?? 'Titre manquant',
      message.notification?.body ?? 'Message manquant',
      notificationDetails,
    );
  }
  Future<void> _storeFcmToken(String fcmToken) async {
    final storage = const FlutterSecureStorage();

    final token = await storage.read(key: Constants.jwtStorageToken);
    try {
      final Uri uri = Uri.parse('${Constants.apiBaseUrl}/api/notifications/fcm_token/$fcmToken');

      await dio.post(
        uri.toString(),
        options: Options(
          headers: {
            'Content-Type': 'application/json',
            "Authorization": "Bearer $token",
          },
        ),
      );
    } catch (error) {
      throw AppException(message: 'Failed to store fcm token');
    }
  }

}
