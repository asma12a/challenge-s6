import 'package:flutter/material.dart';

class CustomDataTable extends StatelessWidget {
  final String title;
  final List<DataColumn> columns;
  final List<DataRow> rows;

  const CustomDataTable({
    super.key,
    required this.title,
    required this.columns,
    required this.rows,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.all(16.0),
      padding: const EdgeInsets.all(16.0),
      decoration: BoxDecoration(
        color: Colors.white, // Fond blanc par défaut
        borderRadius: BorderRadius.circular(8),
        border: Border.all(
          color: Colors.grey[300]!, // Bordure légère grise
          width: 1,
        ),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.stretch,
        children: [
          Text(
            title,
            style: Theme.of(context).textTheme.headlineSmall?.copyWith(
                  fontWeight: FontWeight.bold,
                  color: Theme.of(context).primaryColor,
                ),
          ),
          const SizedBox(height: 16),
          SingleChildScrollView(
            scrollDirection: Axis.horizontal,
            child: DataTable(
              headingRowColor:
                  MaterialStateColor.resolveWith((states) => Colors.grey[200]!),
              columns: columns,
              rows: rows,
            ),
          ),
        ],
      ),
    );
  }
}
