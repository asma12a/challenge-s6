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

    try {
      final response = await dio.get(
        uri,
        options: Options(
          headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer $token',
          },
        ),
      );
      // Vérification de la structure de la réponse
      if (response.data is List) {
        return List<Map<String, dynamic>>.from(response.data);
      } else {
        return [];
      }
    } catch (e) {
      return [];
    }
  }
}
