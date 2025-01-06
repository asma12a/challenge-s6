import 'dart:convert';
import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/models/sport.dart';
import 'package:squad_go/main.dart';

class EventService {
  final storage = const FlutterSecureStorage();

  // GET all events (backoffice)
  Future<List<Map<String, dynamic>>> getEvents() async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    final url = Uri.http(dotenv.env['API_BASE_URL']!, 'api/events');

    try {
      final response = await dio.get(url.toString(),
          options: Options(headers: {
            'Content-Type': 'application/json',
            'Authorization': "Bearer $token",
          }));

      if (response.statusCode != 200) {
        throw Exception('Erreur lors de la récupération des événements.');
      }

      final List<dynamic> data = response.data;
      return List<Map<String, dynamic>>.from(data);
    } catch (error) {
      throw Exception('Erreur: ${error.toString()}');
    }
  }

  Future<List<Map<String, dynamic>>> getSearchResults(
      Map<String, String> params) async {
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
      final response = await dio.get(url.toString(),
          options: Options(headers: {
            'Content-Type': 'application/json',
            'Authorization': "Bearer $token",
          }));

      final List<Map<String, dynamic>> events =
          List<Map<String, dynamic>>.from(json.decode(response.data));
      return events;
    } catch (error) {
      throw AppException(
          message: 'Failed to retrieve events, please try again.');
    }
  }

  // GET event By CODE
  Future<Event> getEventByCode(String code) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    try {
      final uri =
          Uri.http(dotenv.env['API_BASE_URL']!, 'api/events/code/$code');
      final response = await dio.get(
        uri.toString(),
        options: Options(headers: {
          'Content-Type': 'application/json',
          "Authorization": "Bearer $token",
        }),
      );
      final Map<String, dynamic> data = response.data;
      return Event.fromJson(data);
    } catch (error) {
      debugPrint('An error occurred: $error');
      throw AppException(
          message: 'Failed to retrieve event, please try again.');
    }
  }

  // DELETE an event
  Future<void> deleteEvent(String id) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    final url = Uri.http(dotenv.env['API_BASE_URL']!, 'api/events/$id');

    try {
      final response = await dio.delete(url.toString(),
          options: Options(headers: {
            'Content-Type': 'application/json',
            'Authorization': "Bearer $token",
          }));

      if (response.statusCode != 204) {
        throw Exception('Erreur lors de la suppression de l\'événement.');
      }
    } catch (error) {
      throw Exception('Erreur: ${error.toString()}');
    }
  }

  Future<Event> getEventById(String id) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);

    final Uri url = Uri.http(dotenv.env['API_BASE_URL']!, 'api/events/$id');

    try {
      final response = await dio.get(url.toString(),
          options: Options(headers: {
            'Content-Type': 'application/json',
            'Authorization': "Bearer $token",
          }));

      final Map<String, dynamic> event =
          Map<String, dynamic>.from(response.data);
      return Event.fromJson(event);
    } catch (error) {
      throw AppException(
          message: 'Failed to retrieve event, please try again.');
    }
  }

  Future<List<Sport>> getSports() async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    try {
      final uri = Uri.http(dotenv.env['API_BASE_URL']!, 'api/sports');
      final response = await dio.get(uri.toString(),
          options: Options(headers: {
            'Content-Type': 'application/json',
            "Authorization": "Bearer $token",
          }));
      return List<Sport>.from(
          response.data.map((sport) => Sport.fromJson(sport)));
    } catch (error) {
      log.severe('An error occurred while ', {error: error});
      throw AppException(
          message: 'Failed to retrieve sports, please try again.');
    }
  }

  Future<void> createEvent(Map<String, dynamic> event) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    try {
      final uri = Uri.http(dotenv.env['API_BASE_URL']!, 'api/events');
      await dio.post(uri.toString(),
          options: Options(headers: {
            'Content-Type': 'application/json',
            "Authorization": "Bearer $token",
          }),
          data: jsonEncode(event));
    } catch (error) {
      log.severe('An error occurred while ', {error: error});
      throw AppException(message: 'Failed to create event, please try again.');
    }
  }

  Future<void> updateEvent(String id, Map<String, dynamic> event) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    try {
      final uri = Uri.http(dotenv.env['API_BASE_URL']!, 'api/events/$id');
      await dio.put(uri.toString(),
          options: Options(headers: {
            'Content-Type': 'application/json',
            "Authorization": "Bearer $token",
          }),
          data: jsonEncode(event));
    } catch (error) {
      log.severe('An error occurred while ', {error: error});
      throw AppException(message: 'Failed to update event, please try again.');
    }
  }

  Future<List<Event>> getMyEvents() async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    try {
      final uri = Uri.http(dotenv.env['API_BASE_URL']!, 'api/events/user');
      final response = await dio.get(uri.toString(),
          options: Options(headers: {
            'Content-Type': 'application/json',
            "Authorization": "Bearer $token",
          }));
      final List<dynamic> data = response.data;
      return data.map((event) => Event.fromJson(event)).toList();
    } catch (error) {
      log.severe('An error occurred while ', {error: error});
      throw AppException(
          message: 'Failed to retrieve my events, please try again.');
    }
  }

  Future<List<Event>> getRecommendedEvents(
      {double? latitude, double? longitude}) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    try {
      final Map<String, String> queryParams = {};
      if (latitude != null && longitude != null) {
        queryParams['latitude'] = latitude.toString();
        queryParams['longitude'] = longitude.toString();
      }
      final uri = Uri.http(
          dotenv.env['API_BASE_URL']!, 'api/events/recommended', queryParams);
      final response = await dio.get(uri.toString(),
          options: Options(headers: {
            'Content-Type': 'application/json',
            "Authorization": "Bearer $token",
          }));
      final List<dynamic> data = response.data;
      return data.map((event) => Event.fromJson(event)).toList();
    } catch (error) {
      log.severe('An error occurred while ', {error: error});
      throw AppException(
          message: 'Failed to retrieve recommended events, please try again.');
    }
  }
}
