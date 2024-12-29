import 'dart:convert';
import 'dart:developer';

import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/core/models/user_app.dart';

class AuthService {
  final _storage = const FlutterSecureStorage();

  Future<Map<String, dynamic>> signIn(body) async {
    try {
     final uri = Uri.parse('http://localhost:3001/api/auth/login');

      final response = await http.post(uri,
          headers: {
            'Content-Type': 'application/json',
          },
          body: jsonEncode(body));

      final data = jsonDecode(utf8.decode(response.bodyBytes));
      await _storage.write(key: 'squadgo-jwt', value: data['token']);

      return data;
    } catch (error) {
      log('An error occurred while ', error: error);
      throw AppException(message: 'Failed to log in, please try again.');
    }
  }

  static Future<Map<String, dynamic>?> signUp(body) async {
    try {
      final uri = Uri.http(dotenv.env['API_BASE_URL']!, 'api/auth/signup');
      final response = await http.post(
        uri,
        headers: {
          'Content-Type': 'application/json',
        },
        body: jsonEncode(body),
      );
      if (response.statusCode == 201) {
        return null;
      }
      final data = jsonDecode(utf8.decode(response.bodyBytes));
      return data;
    } catch (error) {
      log('An error occurred during sign-up', error: error);
      throw AppException(message: 'Failed to sign up, please try again.');
    }
  }

  Future<UserApp?> getUserInfo() async {
    try {
      final token = await _storage.read(key: 'squadgo-jwt');
      if (token == null) return null;

      final uri = Uri.http(dotenv.env['API_BASE_URL']!, 'api/auth/me');
      final response = await http.get(uri, headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token',
      });

      if (response.statusCode == 200) {
        final data = jsonDecode(utf8.decode(response.bodyBytes));
        return UserApp.fromJson({...data, 'apiToken': token});
      }
    } catch (error) {
      log(
        'An error occurred while fetching user info, error: $error',
      );
      throw AppException(
        message: 'Failed to fetch user info, please try again.',
      );
    }
    return null;
  }
}
