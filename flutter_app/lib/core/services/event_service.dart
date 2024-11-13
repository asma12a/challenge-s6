import 'dart:convert';
import 'dart:developer';
import 'package:flutter_app/core/exceptions/app_exception.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;

import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class EventService {
  static Future<List<Map<String, dynamic>>> getSearchResults(
      Map<String, String> params) async {
    final storage = const FlutterSecureStorage();
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);

    final Uri baseUrl =
        Uri.http(dotenv.env['API_BASE_URL']!, 'api/events/search');

    final Map<String, String> queryParams = {};
    if (params.isNotEmpty) {
      params.forEach((key, value) {
        queryParams[key] = value;
      });
    }

    final Uri url = queryParams.isNotEmpty
        ? baseUrl.replace(queryParameters: queryParams)
        : baseUrl;

    try {
      final response = await http.get(url, headers: {
        'Content-Type': 'application/json',
        'Authorization': "Bearer $token",
      });

      final List<Map<String, dynamic>> events =
          List<Map<String, dynamic>>.from(json.decode(response.body));
      return events;
    } catch (error) {
      throw AppException(
          message: 'Failed to retrieve events, please try again.');
    }
  }

  static Future<List<Map<String, dynamic>>> getSports() async {
    final storage = const FlutterSecureStorage();
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    try {
      final uri = Uri.http(dotenv.env['API_BASE_URL']!, 'api/sports');
      final response = await http.get(
        uri,
        headers: {
          'Content-Type': 'application/json',
          "Authorization": "Bearer $token",
        },
      );
      final data = jsonDecode(utf8.decode(response.bodyBytes));
      return List<Map<String, dynamic>>.from(data);
    } catch (error) {
      log('An error occurred while ', error: error);
      throw AppException(
          message: 'Failed to retrieve sports, please try again.');
    }
  }
}
