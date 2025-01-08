import 'dart:convert';
import 'package:dio/dio.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:http/http.dart' as http;
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/core/models/sport.dart';
import 'package:squad_go/main.dart';


class SportService {
  final storage = const FlutterSecureStorage();

  // GET all sports
  static Future<List<Map<String, dynamic>>> getSports() async {
    final storage = const FlutterSecureStorage();
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);

    final Uri url = Uri.http(dotenv.env['API_BASE_URL']!, 'api/sports');

    try {
      final response = await http.get(url, headers: {
        'Content-Type': 'application/json',
        'Authorization': "Bearer $token",
      });

      if (response.statusCode != 200) {
        throw Exception('Erreur lors de la récupération des sports.');
      }

      final List<dynamic> data = json.decode(response.body);
      return List<Map<String, dynamic>>.from(data);
    } catch (error) {
      throw Exception('Erreur: ${error.toString()}');
    }
  }

  // DELETE a specific sport
  static Future<void> deleteSport(String id) async {
    final storage = const FlutterSecureStorage();
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);

    final Uri url = Uri.http(dotenv.env['API_BASE_URL']!, 'api/sports/$id');

    try {
      final response = await http.delete(url, headers: {
        'Content-Type': 'application/json',
        'Authorization': "Bearer $token",
      });

      if (response.statusCode != 204) {
        throw Exception('Erreur lors de la suppression du sport.');
      }
    } catch (error) {
      throw Exception('Erreur: ${error.toString()}');
    }
  }

  // CREATE a new sport
  static Future<void> createSport(Map<String, dynamic> sportData) async {
    final storage = const FlutterSecureStorage();
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    final Uri url = Uri.http(dotenv.env['API_BASE_URL']!, 'api/sports');

    try {
      final response = await http.post(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': "Bearer $token",
        },
        body: json.encode(sportData),
      );

      if (response.statusCode != 201) {
        throw Exception('Erreur lors de la création du sport.');
      }
    } catch (error) {
      throw Exception('Erreur: ${error.toString()}');
    }
  }

  // UPDATE an existing sport
  static Future<void> updateSport(
      String id, Map<String, dynamic> updates) async {
    final storage = const FlutterSecureStorage();
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    final Uri url = Uri.http(dotenv.env['API_BASE_URL']!, 'api/sports/$id');

    try {
      final response = await http.put(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': "Bearer $token",
        },
        body: json.encode(updates),
      );

      if (response.statusCode != 200) {
        throw Exception(
            'Erreur lors de la mise à jour du sport : ${response.body}');
      }
    } catch (error) {
      throw Exception('Erreur: ${error.toString()}');
    }
  }

  Future<List<Sport>> getUserSports() async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    try {
      final uri = Uri.http(dotenv.env['API_BASE_URL']!, 'api/sports/user');
      final response = await dio.get(uri.toString(),
          options: Options(headers: {
            'Content-Type': 'application/json',
            "Authorization": "Bearer $token",
          }));
      final List<dynamic> data = response.data;
      return data.map((sport) => Sport.fromJson(sport)).toList();
    } catch (error) {
      log.severe('An error occurred while ', {error: error});
      throw AppException(
          message: 'Failed to retrieve my sports, please try again.');
    }
  }
}
