import 'dart:convert';
import 'dart:developer';
import 'package:flutter/cupertino.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/core/models/sport_stat_labels.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/core/models/user_stats.dart';

class SportStatLabelsService {
  final storage = const FlutterSecureStorage();

  Future<List<SportStatLabels>> getStatLabelsBySport(String sportId) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    try {
      final uri = Uri.http(
          dotenv.env['API_BASE_URL']!, 'api/sportstatlabels/$sportId/labels');
      final response = await http.get(
        uri,
        headers: {
          'Content-Type': 'application/json',
          "Authorization": "Bearer $token",
        },
      );
      final data = jsonDecode(utf8.decode(response.bodyBytes));
      return List<SportStatLabels>.from(
        data.map((json) => SportStatLabels.fromJson(json)),
      );
    } catch (error) {
      throw AppException(
          message: 'Failed to retrieve sport stat labels, please try again.');
    }
  }

  Future<List<UserStats>> getUserStatByEvent(String eventId, userId) async {
    debugPrint('UserId $userId');
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    try {
      final uri = Uri.http(dotenv.env['API_BASE_URL']!,
          'api/sportstatlabels/$eventId/$userId/stats');
      final response = await http.get(
        uri,
        headers: {
          'Content-Type': 'application/json',
          "Authorization": "Bearer $token",
        },
      );
      final data = jsonDecode(utf8.decode(response.bodyBytes));
      return List<UserStats>.from(
        data.map((json) => UserStats.fromJson(json)),
      );
    } catch (error) {
      throw AppException(
          message: 'Failed to retrieve sport stat labels, please try again.');
    }
  }

  Future<UserPerformance> getUserPerformanceBySport(String sportId, userId) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    try {
      final uri = Uri.http(dotenv.env['API_BASE_URL']!, 'api/sportstatlabels/$sportId/$userId/performance');
      final response = await http.get(
          uri,
          headers: {
          'Content-Type': 'application/json',
          "Authorization": "Bearer $token",
        },
      );
      final data = jsonDecode(utf8.decode(response.bodyBytes));
      return UserPerformance.fromJson(data);
    } catch (error) {
      throw AppException(
          message: 'Failed to retrieve user performance, please try again.');
    }
  }

  Future<void> addUserStat(Map<String, dynamic> stats, String eventId) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    try {
      final uri = Uri.http(dotenv.env['API_BASE_URL']!,
          'api/sportstatlabels/$eventId/addUserStat');
      await http.post(uri,
          headers: {
            'Content-Type': 'application/json',
            "Authorization": "Bearer $token",
          },
          body: jsonEncode(stats));
    } catch (error) {
      throw AppException(
          message: 'Failed to add user stats, please try again.');
    }
  }

  Future<void> updateUserStat(
      Map<String, dynamic> stats, String eventId) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    try {
      final uri = Uri.http(dotenv.env['API_BASE_URL']!,
          'api/sportstatlabels/$eventId/updateUserStats');
      await http.put(uri,
          headers: {
            'Content-Type': 'application/json',
            "Authorization": "Bearer $token",
          },
          body: jsonEncode(stats));
    } catch (error) {
      throw AppException(message: 'Failed to update user stat, please try again.');
    }
  }

  Future<List<SportStatLabels>> getAllStatLabels() async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    try {
      final uri = Uri.http(dotenv.env['API_BASE_URL']!, 'api/sportstatlabels');
      final response = await http.get(
        uri,
        headers: {
          'Content-Type': 'application/json',
          "Authorization": "Bearer $token",
        },
      );

      if (response.statusCode == 200) {
        final data = jsonDecode(utf8.decode(response.bodyBytes));
        return List<SportStatLabels>.from(
          data.map((json) => SportStatLabels.fromJson(json)),
        );
      } else {
        throw AppException(message: 'Failed to load stat labels');
      }
    } catch (error) {
      log('An error occurred while fetching sport stat labels', error: error);
      throw AppException(
        message:
            'An error occurred while fetching stat labels, please try again.',
      );
    }
  }

  Future<void> createStatLabel(Map<String, dynamic> statLabelData) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    try {
      final uri = Uri.http(dotenv.env['API_BASE_URL']!, 'api/sportstatlabels');
      final response = await http.post(
        uri,
        headers: {
          'Content-Type': 'application/json',
          "Authorization": "Bearer $token",
        },
        body: jsonEncode(statLabelData),
      );

      if (response.statusCode == 201) {
        return;
      } else {
        throw AppException(message: 'Failed to create stat label');
      }
    } catch (error) {
      log('An error occurred while creating stat label', error: error);
      throw AppException(message: 'Failed to create stat label');
    }
  }

  Future<void> updateStatLabel(
      Map<String, dynamic> statLabelData, String statLabelId) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    try {
      final uri = Uri.http(
          dotenv.env['API_BASE_URL']!, 'api/sportstatlabels/$statLabelId');
      final response = await http.put(
        uri,
        headers: {
          'Content-Type': 'application/json',
          "Authorization": "Bearer $token",
        },
        body: jsonEncode(statLabelData),
      );

      if (response.statusCode == 200) {
        return;
      } else {
        throw AppException(message: 'Failed to update stat label');
      }
    } catch (error) {
      log('An error occurred while updating stat label', error: error);
      throw AppException(message: 'Failed to update stat label');
    }
  }

  Future<void> deleteStatLabel(String? statLabelId) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    try {
      final uri = Uri.http(
          dotenv.env['API_BASE_URL']!, 'api/sportstatlabels/$statLabelId');
      final response = await http.delete(
        uri,
        headers: {
          'Content-Type': 'application/json',
          "Authorization": "Bearer $token",
        },
      );

      if (response.statusCode != 204) {
                throw Exception('Erreur lors de la suppression de la statistique.');

      }
    } catch (error) {
      throw Exception('Erreur: ${error.toString()}');
    }
  }
}
