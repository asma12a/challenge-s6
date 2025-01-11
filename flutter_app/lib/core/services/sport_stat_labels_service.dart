import 'dart:convert';
import 'package:dio/dio.dart';
import 'package:flutter/cupertino.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/core/models/sport_stat_labels.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/core/models/user_stats.dart';
import 'package:squad_go/core/utils/constants.dart';
import 'package:squad_go/main.dart';

class SportStatLabelsService {
  final storage = const FlutterSecureStorage();

  Future<List<SportStatLabels>> getStatLabelsBySport(String sportId) async {
    final token = await storage.read(key: Constants.jwtStorageToken);
    try {
      final Uri uri = Uri.parse(
          '${Constants.apiBaseUrl}/api/sportstatlabels/$sportId/labels');

      final response = await dio.get(
        uri.toString(),
        options: Options(
          headers: {
            'Content-Type': 'application/json',
            "Authorization": "Bearer $token",
          },
        ),
      );

      final List<dynamic> sportStatLabels = response.data;
      return sportStatLabels.map((sportStat) => SportStatLabels.fromJson(sportStat)).toList();
    } catch (error) {
      throw AppException(
          message: 'Failed to retrieve sport stat labels, please try again.');
    }
  }

  Future<List<UserStats>> getUserStatByEvent(String eventId, userId) async {
    final token = await storage.read(key: Constants.jwtStorageToken);
    try {
      final Uri uri = Uri.parse(
          '${Constants.apiBaseUrl}/api/sportstatlabels/$eventId/$userId/stats');

      final response = await dio.get(
        uri.toString(),
        options: Options(
          headers: {
            'Content-Type': 'application/json',
            "Authorization": "Bearer $token",
          },
        ),
      );

      final List<dynamic> userStats = response.data;
      return userStats.map((userStat) => UserStats.fromJson(userStat)).toList();
    } catch (error) {
      throw AppException(
          message: 'Failed to retrieve sport stat labels, please try again.');
    }
  }

  Future<UserPerformance> getUserPerformanceBySport(
      String sportId, userId) async {
    final token = await storage.read(key: Constants.jwtStorageToken);
    try {
      final Uri uri = Uri.parse(
          '${Constants.apiBaseUrl}/api/sportstatlabels/$sportId/$userId/performance');

      final response = await dio.get(
        uri.toString(),
        options: Options(
          headers: {
            'Content-Type': 'application/json',
            "Authorization": "Bearer $token",
          },
        ),
      );

      final Map<String, dynamic> data = response.data;
      return UserPerformance.fromJson(data);

    } catch (error) {
      throw AppException(
          message: 'Failed to retrieve user performance, please try again.');
    }
  }

  Future<void> addUserStat(Map<String, dynamic> stats, String eventId) async {
    final token = await storage.read(key: Constants.jwtStorageToken);
    try {
      final Uri uri = Uri.parse(
          '${Constants.apiBaseUrl}/api/sportstatlabels/$eventId/addUserStat');

      await dio.post(
        uri.toString(),
        options: Options(
          headers: {
            'Content-Type': 'application/json',
            "Authorization": "Bearer $token",
          },
        ),
        data: jsonEncode(stats),
      );
    } catch (error) {
      throw AppException(
          message: 'Failed to add user stats, please try again.');
    }
  }

  Future<void> updateUserStat(
      Map<String, dynamic> stats, String eventId) async {
    final token = await storage.read(key: Constants.jwtStorageToken);
    try {
      final Uri uri = Uri.parse(
          '${Constants.apiBaseUrl}/api/sportstatlabels/$eventId/updateUserStats');

      await dio.put(
        uri.toString(),
        options: Options(
          headers: {
            'Content-Type': 'application/json',
            "Authorization": "Bearer $token",
          },
        ),
        data: jsonEncode(stats),
      );
    } catch (error) {
      throw AppException(
          message: 'Failed to update user stat, please try again.');
    }
  }

  Future<List<SportStatLabels>> getAllStatLabels() async {
    final token = await storage.read(key: Constants.jwtStorageToken);
    try {
      final Uri uri = Uri.parse('${Constants.apiBaseUrl}/api/sportstatlabels');

      final response = await dio.get(
        uri.toString(),
        options: Options(
          headers: {
            'Content-Type': 'application/json',
            "Authorization": "Bearer $token",
          },
        ),
      );

      final data = jsonDecode(utf8.decode(response.data));
      return List<SportStatLabels>.from(
        data.map((json) => SportStatLabels.fromJson(json)),
      );
    } catch (error) {
      throw AppException(
        message:
            'An error occurred while fetching stat labels, please try again.',
      );
    }
  }

  Future<void> createStatLabel(Map<String, dynamic> statLabelData) async {
    final token = await storage.read(key: Constants.jwtStorageToken);
    try {
      final Uri uri = Uri.parse('${Constants.apiBaseUrl}/api/sportstatlabels');

      await dio.post(
        uri.toString(),
        options: Options(
          headers: {
            'Content-Type': 'application/json',
            "Authorization": "Bearer $token",
          },
        ),
        data: jsonEncode(statLabelData),
      );
    } catch (error) {
      throw AppException(message: 'Failed to create stat label');
    }
  }

  Future<void> updateStatLabel(
      Map<String, dynamic> statLabelData, String statLabelId) async {
    final token = await storage.read(key: Constants.jwtStorageToken);
    try {
      final Uri uri =
          Uri.parse('${Constants.apiBaseUrl}/api/sportstatlabels/$statLabelId');

      await dio.put(
        uri.toString(),
        options: Options(
          headers: {
            'Content-Type': 'application/json',
            "Authorization": "Bearer $token",
          },
        ),
        data: jsonEncode(statLabelData),
      );
    } catch (error) {
      throw AppException(message: 'Failed to update stat label');
    }
  }

  Future<void> deleteStatLabel(String? statLabelId) async {
    final token = await storage.read(key: Constants.jwtStorageToken);
    try {
      final Uri uri =
          Uri.parse('${Constants.apiBaseUrl}/api/sportstatlabels/$statLabelId');

      final response = await dio.delete(
        uri.toString(),
        options: Options(
          headers: {
            'Content-Type': 'application/json',
            "Authorization": "Bearer $token",
          },
        ),
      );

      if (response.statusCode != 204) {
        throw AppException(message: 'Failed to delete stat label.');
      }
    } catch (error) {
      throw AppException(message: 'Error: ${error.toString()}');
    }
  }
}
