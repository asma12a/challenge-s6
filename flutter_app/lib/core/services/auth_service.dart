import 'dart:convert';
import 'dart:developer';

import 'package:flutter_app/core/exceptions/app_exception.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;

class AuthService {
  static Future<Map<String, dynamic>> signIn(body) async {
    try {
      final uri = Uri.http(dotenv.env['API_BASE_URL']!, 'api/auth/login');
      final response = await http.post(
          uri,
          headers: {
            'Content-Type': 'application/json',
          },
          body: jsonEncode(body)
      );
      final data = jsonDecode(utf8.decode(response.bodyBytes));
      print(data);
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
  }}
