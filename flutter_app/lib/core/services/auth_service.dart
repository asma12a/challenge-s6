import 'package:dio/dio.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/core/models/user_app.dart';
import 'package:squad_go/main.dart';

class AuthService {
  final _storage = const FlutterSecureStorage();

  Future<Map<String, dynamic>> signIn(body) async {
          print('DEBUT signIn ===');
      print('signIn  ${dotenv.env['API_BASE_URL']!}');

    try {
      final uri = Uri.http(dotenv.env['API_BASE_URL']!, 'api/auth/login');
      print('signIn on service ${dotenv.env['API_BASE_URL']!}');

      final response = await dio.post(
        uri.toString(),
        options: Options(
          headers: {
            'Content-Type': 'application/json',
          },
        ),
        data: body,
      );

      print('response $response');

      final data = response.data;
      print('data before token $data');

      await _storage.write(
          key: dotenv.env['JWT_STORAGE_KEY']!, value: data['token']);
      print('data after token $data');

      return data;
    } catch (error) {
      log.severe('An error occurred while ', {error: error});
      throw AppException(message: 'Failed to log in, please try again.');
    }
  }

  Future<Map<String, dynamic>?> signUp(body) async {
    try {
      final uri = Uri.http(dotenv.env['API_BASE_URL']!, 'api/auth/signup');
      final response = await dio.post(
        uri.toString(),
        options: Options(
          headers: {
            'Content-Type': 'application/json',
          },
        ),
        data: body,
      );

      if (response.statusCode == 201) {
        return null;
      }

      return response.data;
    } catch (error) {
      log.severe('An error occurred while ', {error: error});
      throw AppException(message: 'Failed to sign up, please try again.');
    }
  }

  Future<UserApp?> getUserInfo() async {
    try {
      final token = await _storage.read(key: 'squadgo-jwt');
      if (token == null) return null;

      final uri = Uri.http(dotenv.env['API_BASE_URL']!, 'api/auth/me');
      final response = await dio.get(
        uri.toString(),
        options: Options(
          headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer $token',
          },
        ),
      );

      if (response.statusCode == 200) {
        return UserApp.fromJson({...response.data, 'apiToken': token});
      }
    } catch (error) {
      log.severe(
        'An error occurred while fetching user info, error: $error',
      );
      throw AppException(
        message: 'Failed to fetch user info, please try again.',
      );
    }
    return null;
  }
}
