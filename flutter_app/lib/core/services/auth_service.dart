import 'package:dio/dio.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/core/models/user_app.dart';
import 'package:squad_go/core/utils/constants.dart';
import 'package:squad_go/main.dart';

class AuthService {
  final _storage = const FlutterSecureStorage();

  Future<Map<String, dynamic>> signIn(body) async {
    try {
      final uri = '${Constants.apiBaseUrl}/api/auth/login';

      final response = await dio.post(
        uri,
        options: Options(
          headers: {
            'Content-Type': 'application/json',
          },
        ),
        data: body,
      );

      final data = response.data;

      await _storage.write(
          key: Constants.jwtStorageToken, value: data['token']);
      return data;
    } catch (error) {
      log.severe('An error occurred while ', {error: error});
      throw AppException(message: 'Failed to log in, please try again.');
    }
  }

  Future<Map<String, dynamic>?> signUp(body) async {
    try {
      final uri = '${Constants.apiBaseUrl}/api/auth/signup';

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
      final token = await _storage.read(key: Constants.jwtStorageToken);
      if (token == null) return null;

      final uri = '${Constants.apiBaseUrl}/api/auth/me';
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
