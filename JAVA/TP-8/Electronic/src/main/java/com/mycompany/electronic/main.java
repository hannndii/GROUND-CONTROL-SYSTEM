/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 */

package com.mycompany.electronic;

/**
 *
 * @author tupog
 */
public class main {
    public static void main(String[] args) {
        Appliance tv = new Television(32); 
        Appliance audio = new Audio("Black"); 

        tv.turnOn();
        tv.turnOff();

        audio.turnOn();
        audio.turnOff();
    }
}
