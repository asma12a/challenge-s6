import 'package:dio/dio.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/main.dart';

class TeamService {
  final storage = const FlutterSecureStorage();

  Future<void> joinTeam(String eventID, String teamID) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);

    final Uri url = Uri.http(
        dotenv.env['API_BASE_URL']!, 'api/events/$eventID/teams/$teamID/join');

    try {
      await dio.post(url.toString(),
          options: Options(headers: {
            'Content-Type': 'application/json',
            'Authorization': "Bearer $token",
          }));
    } catch (error) {
      throw AppException(message: 'Failed to join team, please try again.');
    }
  }
}
