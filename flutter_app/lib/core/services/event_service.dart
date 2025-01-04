import 'dart:convert';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:http/http.dart' as http;
import 'dart:developer';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

class EventService {

  // GET event By CODE
  static Future<Map<String, dynamic>> getEventByCode(String code) async {
    final storage = const FlutterSecureStorage();
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    try {
      final uri =
          Uri.http(dotenv.env['API_BASE_URL']!, 'api/events/code/$code');
      final response = await http.get(
        uri,
        headers: {
          'Content-Type': 'application/json',
          "Authorization": "Bearer $token",
        },
      );
      final data = jsonDecode(utf8.decode(response.bodyBytes));
      return data;
    } catch (error) {
      log('An error occurred while ', error: error);
      throw AppException(
          message: 'Failed to retrieve event, please try again.');
    }
  }

  // GET all events
  static Future<List<Map<String, dynamic>>> getEvents() async {
    final storage = const FlutterSecureStorage();
    // final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    //final url = Uri.http(dotenv.env['API_BASE_URL']!, 'api/events');

    final token = await storage.read(key: 'squadgo-jwt');

    final url = Uri.parse('http://localhost:3001/api/events');
    try {
      final response = await http.get(url, headers: {
        'Content-Type': 'application/json',
        'Authorization': "Bearer $token",
      });

      if (response.statusCode != 200) {
        throw Exception('Erreur lors de la récupération des événements.');
      }

      final List<dynamic> data = json.decode(response.body);
      return List<Map<String, dynamic>>.from(data);
    } catch (error) {
      throw Exception('Erreur: ${error.toString()}');
    }
  }

// CREATE an event
  static Future<void> createEvent(Map<String, dynamic> event) async {
    final storage = const FlutterSecureStorage();
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    try {
      final uri = Uri.http(dotenv.env['API_BASE_URL']!, 'api/events');
      await http.post(uri,
          headers: {
            'Content-Type': 'application/json',
            "Authorization": "Bearer $token",
          },
          body: jsonEncode(event));
    } catch (error) {
      log('An error occurred while ', error: error);
      throw AppException(message: 'Failed to create event, please try again.');
    }
  }

  // DELETE an event
  static Future<void> deleteEvent(String id) async {
    final storage = const FlutterSecureStorage();
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    final url = Uri.http(dotenv.env['API_BASE_URL']!, 'api/events/$id');

    try {
      final response = await http.delete(url, headers: {
        'Content-Type': 'application/json',
        'Authorization': "Bearer $token",
      });

      if (response.statusCode != 204) {
        throw Exception('Erreur lors de la suppression de l\'événement.');
      }
    } catch (error) {
      throw Exception('Erreur: ${error.toString()}');
    }
  }

  // UPDATE an existing event
  static Future<void> updateEvent(
      String id, Map<String, dynamic> updates) async {
    final storage = const FlutterSecureStorage();
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);

    try {
      final url = Uri.http(dotenv.env['API_BASE_URL']!, 'api/events/$id');
      final response = await http.put(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': "Bearer $token",
        },
        body: json.encode(updates),
      );

      if (response.statusCode != 200) {
        throw Exception('Erreur lors de la mise à jour de l\'événement.');
      }
    } catch (error) {
      throw Exception('Erreur: ${error.toString()}');
    }
  }

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
