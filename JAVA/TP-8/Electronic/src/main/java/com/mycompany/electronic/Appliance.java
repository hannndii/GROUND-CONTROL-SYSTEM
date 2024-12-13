/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.electronic;

/**
 *
 * @author tupog
 */
class Appliance {
    protected String product;
    protected String place;

    public void turnOn() {
        System.out.println(product + " : On");
    }

    public void turnOff() {
        System.out.println(product + " : Off");
    }

    public void setProduct(String product) {
        this.product = product;
    }
}
