import 'dart:convert';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/core/models/user_app.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;

class UserService {
  // GET all users
  static Future<List<UserApp>> getUsers() async {
    final storage = const FlutterSecureStorage();
    //     final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);

    final token = await storage.read(key: 'squadgo-jwt');

    // final Uri url = Uri.http(dotenv.env['API_BASE_URL']!, 'api/users');
    final url = Uri.parse('http://localhost:3001/api/users');

    try {
      final response = await http.get(url, headers: {
        'Content-Type': 'application/json',
        'Authorization': "Bearer $token",
      });

      if (response.statusCode != 200) {
        throw Exception('Erreur lors de la récupération des utilisateurs.');
      }

      final List<dynamic> data = json.decode(response.body);
      return data.map((user) => UserApp.fromJson(user)).toList();
    } catch (error) {
      throw Exception('Erreur: ${error.toString()}');
    }
  }

  // DELETE a specific user
  static Future<void> deleteUser(String id) async {
    final storage = const FlutterSecureStorage();
    // final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    final token = await storage.read(key: 'squadgo-jwt');

    // final Uri url = Uri.http(dotenv.env['API_BASE_URL']!, 'api/users/$id');
    final url = Uri.parse('http://localhost:3001/api/users/$id');

    try {
      final response = await http.delete(url, headers: {
        'Content-Type': 'application/json',
        'Authorization': "Bearer $token",
      });

      if (response.statusCode != 200) {
        throw Exception('Erreur lors de la suppression de l\'utilisateur.');
      }
    } catch (error) {
      throw Exception('Erreur: ${error.toString()}');
    }
  }

  // CREATE a new user
  static Future<void> createUser(Map<String, dynamic> userData) async {
    final storage = const FlutterSecureStorage();
    // final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    final token = await storage.read(key: 'squadgo-jwt');

    // final Uri url = Uri.http(dotenv.env['API_BASE_URL']!, 'api/users');
    final url = Uri.parse('http://localhost:3001/api/users');

    try {
      final response = await http.post(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': "Bearer $token",
        },
        body: json.encode(userData),
      );

      print('body ================= $userData');

      if (response.statusCode != 201) {
        throw Exception('Erreur lors de la création de l\'utilisateur.');
      }
    } catch (error) {
      throw Exception('Erreur: ${error.toString()}');
    }
  }

  // UPDATE an existing user
  static Future<void> updateUser(
      String id, Map<String, dynamic> updates) async {
    final storage = const FlutterSecureStorage();
    // final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    final token = await storage.read(key: 'squadgo-jwt');

    // final Uri url = Uri.http(dotenv.env['API_BASE_URL']!, 'api/users/$id');
    final url = Uri.parse('http://localhost:3001/api/users/$id');

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
            'Erreur lors de la mise à jour de l\'utilisateur : ${response.body}');
      }
    } catch (error) {
      throw Exception('Erreur: ${error.toString()}');
    }
  }
}
