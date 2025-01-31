import 'package:dio/dio.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/main.dart';

const apiBaseUrl = String.fromEnvironment('API_BASE_URL');
const jwtStorageToken = String.fromEnvironment('JWT_STORAGE_KEY');

class LogService {
  static Future<List<Map<String, dynamic>>> getLogs() async {
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

    final List<dynamic> data = response.data;
    return List<Map<String, dynamic>>.from(data);
  }
}
