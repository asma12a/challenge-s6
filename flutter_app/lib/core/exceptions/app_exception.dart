class AppException implements Exception {
  final String message;

  AppException({required this.message});

  factory AppException.from(dynamic error) {
    if (error is AppException) {
      return error;
    }

    return AppException(message: 'An unknown error occurred');
  }
}