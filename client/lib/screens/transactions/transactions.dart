import 'package:flutter/material.dart';
import 'package:flutter_form_builder/flutter_form_builder.dart';

part 'widgets/_body.dart';

class TransactionsScreen extends StatelessWidget {
  const TransactionsScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Wallet'),
      ),
      body: const _Body(),
    );
  }
}
