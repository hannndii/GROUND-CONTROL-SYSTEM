/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 */

package com.mycompany.jurnal_08;

/**
 *
 * @author tupog
 */
public class VehicleTest {
    public static void main(String[] args) {
        Vehicle[] vehicles = {
            new Truck("B1234CD", 10000, 0.2, 500),
            new Van("D5678EF", 2000, 10.0, 15),
            new Motorcycle("D5678EF", 2000, 10.0, 15),
            new HeavyDutyTruck("D5678EF", 2000, 10.0, 15)
        };
        for (Vehicle vehicle : vehicles) {
            vehicle.displayInfo(); // Polymorphism dan dynamic binding digunakan di sini
        }
    }
}
