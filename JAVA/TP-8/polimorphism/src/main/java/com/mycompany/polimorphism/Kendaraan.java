/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.polimorphism;

/**
 *
 * @author tupog
 */
public class Kendaraan {
    protected String namaKendaraan;
    public Kendaraan(String namaKendaraan) {
        this.namaKendaraan = namaKendaraan;
    }
    public String getInfo(){
        return "Kendaraan: " + namaKendaraan;
    }
}