import 'package:client/blocs/transaction/transaction_bloc.dart';
import 'package:client/blocs/wallet/wallet_bloc.dart';
import 'package:client/models/transaction.dart';
import 'package:client/screens/wallet/wallet.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_form_builder/flutter_form_builder.dart';

part 'widgets/_body.dart';

class TransactionsScreen extends StatelessWidget {
  const TransactionsScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Colors.white,
        elevation: 0,
        title: const Text(
          'Wallet',
          style: TextStyle(color: Colors.black),
        ),
        actions: [
          IconButton(
            onPressed: () {
              Navigator.of(context).push(
                MaterialPageRoute(builder: (_) => const WalletScreen())
              );
            },
            icon: const Icon(
              CupertinoIcons.paperplane_fill,
              color: Colors.black,
            ),
          ),
        ],
      ),
      body: const _Body(),
    );
  }
}
