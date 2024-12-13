/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.soal_tp_5;

/**
 *
 * @author tupog
 */
abstract class Account {
    private final String number;
    private final String name;
    private final int balance;

    // Constructor
    public Account(String number, String name, int balance) {
        this.number = number;
        this.name = name;
        this.balance = balance;
    }

    // Getter methods
    public String getNumber() {
        return number;
    }

    public String getName() {
        return name;
    }

    public int getBalance() {
        return balance;
    }

    // Method abstrak openAccount
    public abstract void openAccount();
}