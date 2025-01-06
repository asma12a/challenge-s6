import 'package:json_annotation/json_annotation.dart';

@JsonSerializable()
class UserApp {
  final String id;
  final String name;
  final String email;
  final List<UserRole> roles;
  final String apiToken;

  const UserApp({
    required this.id,
    required this.name,
    required this.email,
    required this.roles,
    required this.apiToken,
  });

  factory UserApp.fromJson(Map<String, dynamic> data) {
    return UserApp(
      id: data['id'] ?? '',
      name: data['name'] ?? '',
      email: data['email'] ?? '',
      roles: (data['roles'] as List<dynamic>)
          .map((role) => UserRole.values.firstWhere((e) => e.name == role))
          .toList(),
      apiToken: data['apiToken'] ?? '',
    );
  }
}

enum UserRole { admin, user }
