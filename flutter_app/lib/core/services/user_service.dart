import 'dart:convert';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/core/models/user_app.dart';
import 'package:http/http.dart' as http;

const apiBaseUrl = String.fromEnvironment('API_BASE_URL');
const jwtStorageToken = String.fromEnvironment('JWT_STORAGE_KEY');

class UserService {
  // GET all users
  static Future<List<UserApp>> getUsers() async {
    final storage = const FlutterSecureStorage();
    final token = await storage.read(key: jwtStorageToken);

    try {
      final Uri url = Uri.parse('$apiBaseUrl/api/users');

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
    final token = await storage.read(key: jwtStorageToken);

    final Uri url = Uri.parse('$apiBaseUrl/api/users/$id');

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
    final token = await storage.read(key: jwtStorageToken);

    final Uri url = Uri.parse('$apiBaseUrl/api/users');

    try {
      final response = await http.post(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': "Bearer $token",
        },
        body: json.encode(userData),
      );

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
    final token = await storage.read(key: jwtStorageToken);

    final Uri url = Uri.parse('$apiBaseUrl/api/users/$id');

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
