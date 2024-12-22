import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/core/models/user_app.dart';
import 'package:squad_go/core/services/auth_service.dart';

class AuthState with ChangeNotifier {
  final AuthService _authService = AuthService();
  final FlutterSecureStorage _storage = const FlutterSecureStorage();

  UserApp? _user;
  UserApp? get userInfo => _user;
  bool get isAuthenticated => _user != null;
  bool get isAdmin => _user?.roles.contains(UserRole.admin) ?? false;

  void setUser(UserApp? user) {
    _user = user;
    notifyListeners();
  }

  Future<Map<String, dynamic>> login(String email, String password) async {
    try {
      final loginData =
          await _authService.signIn({'email': email, 'password': password});

      debugPrint('loginData: $loginData');
      if (loginData['status'] == 'error') {
        return loginData;
      }
      // make UserApp from : loginData['user'] props + loginData['token']
      final userData = UserApp.fromJson({
        ...loginData['user'],
        'apiToken': loginData['token'],
      });
      setUser(userData);

      return loginData;
    } catch (e) {
      log('Login failed: $e');
      throw AppException(message: 'Failed to log in. Please try again.');
    }
  }

  Future<void> logout() async {
    try {
      _user = null;
      await _storage.delete(key: dotenv.env['JWT_STORAGE_KEY']!);
      notifyListeners();
    } catch (e) {
      log('Logout failed: $e');
      throw AppException(message: 'Failed to log out. Please try again.');
    }
  }

  Future<bool> tryLogin() async {
    try {
      final token = dotenv.env['JWT_STORAGE_KEY'] != null
          ? await _storage.read(key: dotenv.env['JWT_STORAGE_KEY']!)
          : null;

      if (token == null) {
        return false;
      }

      final userData = await _authService.getUserInfo();
      if (userData != null) {
        setUser(userData);
        return true;
      }
      return false;
    } catch (e) {
      log('Failed to load user: $e');
      return false;
    }
  }
}