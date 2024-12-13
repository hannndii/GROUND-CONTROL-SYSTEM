/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.polimorphism;

/**
 *
 * @author tupog
 */
public class Motor extends Kendaraan{
    public Motor(String namaKendaraan){
        super(namaKendaraan);
    }
    @Override
    public String getInfo(){
        return "Motor: " + namaKendaraan;
    }
}
