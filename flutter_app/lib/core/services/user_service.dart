import 'dart:convert';
import 'package:dio/dio.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/core/models/user_app.dart';
import 'package:squad_go/core/utils/constants.dart';
import 'package:squad_go/main.dart';

class UserService {
  // GET all users
  static Future<List<UserApp>> getUsers() async {
    final storage = const FlutterSecureStorage();
    final token = await storage.read(key: Constants.jwtStorageToken);

    try {
      final Uri url = Uri.parse('${Constants.apiBaseUrl}/api/users');

      final response = await dio.get(
        url.toString(),
        options: Options(
          headers: {
            'Content-Type': 'application/json',
            'Authorization': "Bearer $token",
            'Cache-Control': "no-cache",
          },
        ),
      );
      final List<dynamic> data = response.data;
      return data.map((user) => UserApp.fromJson(user)).toList();
    } catch (error) {
      throw Exception('Erreur: ${error.toString()}');
    }
  }

  // DELETE a specific user
  static Future<void> deleteUser(String id) async {
    final storage = const FlutterSecureStorage();
    final token = await storage.read(key: Constants.jwtStorageToken);

    final Uri url = Uri.parse('${Constants.apiBaseUrl}/api/users/$id');

    try {
      final response = await dio.delete(
        url.toString(),
        options: Options(
          headers: {
            'Content-Type': 'application/json',
            'Authorization': "Bearer $token",
          },
        ),
      );

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
    final token = await storage.read(key: Constants.jwtStorageToken);

    final Uri url = Uri.parse('${Constants.apiBaseUrl}/api/users');

    try {
      final response = await dio.post(
        url.toString(),
        options: Options(
          headers: {
            'Content-Type': 'application/json',
            'Authorization': "Bearer $token",
          },
        ),
        data: json.encode(userData),
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
    final token = await storage.read(key: Constants.jwtStorageToken);

    final Uri url = Uri.parse('${Constants.apiBaseUrl}/api/users/$id');

    try {
      final response = await dio.put(
        url.toString(),
        options: Options(
          headers: {
            'Content-Type': 'application/json',
            'Authorization': "Bearer $token",
          },
        ),
        data: json.encode(updates),
      );

      if (response.statusCode != 200) {
        throw Exception(
            'Erreur lors de la mise à jour de l\'utilisateur : ${response.data}');
      }
    } catch (error) {
      throw Exception('Erreur: ${error.toString()}');
    }
  }

  static Future<void> updateSelfUser(
      String id, Map<String, dynamic> updates) async {
    final storage = const FlutterSecureStorage();
    final token = await storage.read(key: Constants.jwtStorageToken);

    final Uri url = Uri.parse('${Constants.apiBaseUrl}/api/users/$id/user');

    try {
      final response = await dio.put(
        url.toString(),
        options: Options(
          headers: {
            'Content-Type': 'application/json',
            'Authorization': "Bearer $token",
          },
        ),
        data: json.encode(updates),
      );

      if (response.statusCode != 200) {
        throw Exception(
            'Erreur lors de la mise à jour de l\'utilisateur : ${response.data}');
      }
    } catch (error) {
      throw Exception('Erreur: ${error.toString()}');
    }
  }

  static Future<Map<String, dynamic>?> updateUserPassword(
      String id, Map<String, dynamic> updates) async {
    final storage = const FlutterSecureStorage();
    final token = await storage.read(key: Constants.jwtStorageToken);

    final Uri url = Uri.parse('${Constants.apiBaseUrl}/api/users/$id/password');

    try {
      final response = await dio.put(
        url.toString(),
        options: Options(
          headers: {
            'Content-Type': 'application/json',
            'Authorization': "Bearer $token",
          },
        ),
        data: json.encode(updates),
      );
      if (response.statusCode == 200) {
        return null;
      }
      return response.data;
    } catch (error) {
      log.severe('An error occurred while ', {error: error});
      throw AppException(
          message: 'Failed to update password, please try again.');
    }
  }
}
