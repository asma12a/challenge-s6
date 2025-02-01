import 'package:json_annotation/json_annotation.dart';

class User {
  final String id;
  final String? name;
  final String? email;

  const User({
    required this.id,
    this.name,
    this.email,
  });

  factory User.fromJson(Map<String, dynamic> data) {
    return User(
      id: data['id'] ?? '',
      name: data['name'] ?? '',
      email: data['email'] ?? '',
    );
  }
}

