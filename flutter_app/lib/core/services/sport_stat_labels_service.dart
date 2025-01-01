import 'dart:convert';
import 'dart:developer';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/core/models/sport_stat_labels.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class SportStatLabelsService {

  static Future<List<SportStatLabels>> getStatLabelsBySport(String sportId) async {
    final storage = const FlutterSecureStorage();
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    try {
      final uri = Uri.http(dotenv.env['API_BASE_URL']!, 'api/sportstatlabels/$sportId/labels');
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
      log('An error occurred while ', error: error);
      throw AppException(
          message: 'Failed to retrieve sport stat labels, please try again.');
    }
  }


}
