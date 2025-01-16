import 'package:dio/dio.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/core/models/action_log.dart'; // Assurez-vous d'avoir ce modèle pour ActionLog
import 'dart:convert';

import 'package:squad_go/main.dart';

const apiBaseUrl = String.fromEnvironment('API_BASE_URL');
const jwtStorageToken = String.fromEnvironment('JWT_STORAGE_KEY');

class LogService {
  static Future<List<ActionLog>> getLogs() async {
    final storage = const FlutterSecureStorage();

    final token = await storage.read(key: jwtStorageToken);

    final uri = '$apiBaseUrl/api/actionlogs';
    final response = await dio.get(
      uri.toString(),
      options: Options(
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      ),
    );
    if (response.statusCode == 200) {
      List<dynamic> data =
          response.data is List ? response.data : json.decode(response.data);
      return data.map((log) => ActionLog.fromJson(log)).toList();
    } else {
      throw Exception('Failed to load logs');
    }
  }
}
