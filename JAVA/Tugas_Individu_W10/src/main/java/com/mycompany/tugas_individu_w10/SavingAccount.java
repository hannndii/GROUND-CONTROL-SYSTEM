/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.tugas_individu_w10;

/**
 *
 * @author tupog
 */
public class SavingAccount implements Account {
    private double balance;
    private double interestRate;

    public SavingAccount(double interestRate) {
        this.balance = 0;
        this.interestRate = interestRate;
    }

    @Override
    public void deposit(double amount) {
        balance += amount;
        System.out.println("Deposited to Saving Account: " + amount);
    }

    @Override
    public void withdraw(double amount) {
        if (balance >= amount) {
            balance -= amount;
            System.out.println("Withdrawn from Saving Account: " + amount);
        } else {
            System.out.println("Insufficient balance in Saving Account.");
        }
    }

    public double getBalance() {
        return balance;
    }
}
