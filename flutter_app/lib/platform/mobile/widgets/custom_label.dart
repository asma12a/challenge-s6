import 'package:flutter/material.dart';

class CustomLabel extends StatelessWidget {
  final String label;
  final IconData? icon;
  final Color? color;
  final Color? iconColor;
  final Color? backgroundColor;

  const CustomLabel({
    super.key,
    required this.label,
    this.icon,
    this.color,
    this.iconColor,
    this.backgroundColor,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        color: backgroundColor ?? Colors.grey.shade300,
        borderRadius: BorderRadius.circular(8),
      ),
      padding: EdgeInsets.symmetric(horizontal: 6, vertical: 3),
      child: Row(
        children: [
          if (icon != null)
            Icon(
              icon,
              color: iconColor,
            ),
          if (icon != null) SizedBox(width: 8),
          Text(label, style: TextStyle(color: color)),
        ],
      ),
    );
  }
}
