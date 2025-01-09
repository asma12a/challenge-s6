class ActionLog {
  final String id;
  final String userId;
  final String action;
  final String description;
  final DateTime createdAt;

  ActionLog({
    required this.id,
    required this.userId,
    required this.action,
    required this.description,
    required this.createdAt,
  });

  factory ActionLog.fromJson(Map<String, dynamic> json) {
    return ActionLog(
      id: json['id'].toString(),
      userId: json['user_action_logs'] ?? '', 
      action: json['action'] ?? '',
      description: json['description'] ?? '',
      createdAt: DateTime.parse(json['created_at'] ?? DateTime.now().toString()),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'user_action_logs': userId, 
      'action': action,
      'description': description,
      'created_at': createdAt.toIso8601String(),
    };
  }
}
