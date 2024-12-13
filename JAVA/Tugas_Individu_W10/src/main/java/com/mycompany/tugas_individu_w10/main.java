/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 */

package com.mycompany.tugas_individu_w10;

/**
 *
 * @author tupog
 */
public class main {
    public static void main(String[] args) {
        Account savingAccount = new SavingAccount(0.05); 
        Account currentAccount = new CurrentAccount(500); 

        savingAccount.deposit(1000);
        savingAccount.withdraw(500);
        
        currentAccount.deposit(200);
        currentAccount.withdraw(600);
    }
}