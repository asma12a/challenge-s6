import 'package:flutter/material.dart';
import 'package:squad_go/core/utils/connectivity_handler.dart';

class ConnectivityState with ChangeNotifier {
  bool _isConnected = true;
  bool get isConnected => _isConnected;

  ConnectivityState() {
    ConnectivityHandler().connectionChangeStatusController.stream.listen(
      (isConnected) {
        _isConnected = isConnected;
        notifyListeners();
      },
    );
  }
}
