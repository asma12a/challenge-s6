import 'package:dio/dio.dart';
import 'package:google_sign_in/google_sign_in.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/core/models/user_app.dart';
import 'package:squad_go/core/services/notification_service.dart';
import 'package:squad_go/core/utils/constants.dart';
import 'package:squad_go/main.dart';


class AuthService {
  final _storage = const FlutterSecureStorage();
  final GoogleSignIn _googleSignIn = GoogleSignIn(
    clientId: String.fromEnvironment('GOOGLE_CLIENT_ID'),
    scopes: [
      'email',
      'https://www.googleapis.com/auth/contacts.readonly',
    ],
  );

  // Méthode de connexion Google
  Future<Map<String, dynamic>> signInWithGoogle() async {
    try {
      // Tentative de connexion avec Google
      final GoogleSignInAccount? googleUser = await _googleSignIn.signIn();

      if (googleUser == null) {
        throw AppException(message: 'Google sign-in failed, please try again.');
      }

      GoogleSignInAuthentication googleAuth;
      googleAuth = await googleUser.authentication;

      final uri = 'http://127.0.0.1:3001/auth/google/login';
      Response response;

      response = await dio.get(
        uri,
        options: Options(
          headers: {'Content-Type': 'application/json'},
        ),
        data: {
          'accessToken': googleAuth.accessToken,
          'idToken': googleAuth.idToken,
        },
      );

      final data = response.data;
      print(
          '============== response ==============  ${response.data['token']}');

      final String authUrl = response.data['auth_url'];
      print('Redirection vers l\'URL d\'authentification Google: $authUrl');

      await _storage.write(
          key: Constants.jwtStorageToken, value: googleAuth.idToken);

      return data;
    } on AppException catch (error) {
      print("error details ${error.message}");
      log.severe('An unexpected error occurred', {'error': error.toString()});
      throw AppException(
          message: 'An unexpected error occurred, please try again.');
    }
  }

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

      print(data);

      await _storage.write(
          key: Constants.jwtStorageToken, value: data['token']);
      await NotificationService().initNotifications();
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
      print('=================================== token $token');

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
