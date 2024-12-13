/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 */

package com.mycompany.product;

/**
 *
 * @author tupog
 */
public class Main {

    public static void main(String[] args) {
        Product[] arrProduk = new Product[3]; 
        ProductManager pM = new ProductManager();
        Product produkA = new Product("Laptop", 15000000.0, 10);
        Product produkB = new Product("Mouse", 150000.0, 50);
        Product produkC = new Product("Keyboard", 300000.0, 30);
        
        arrProduk[0] = produkA;
        arrProduk[1] = produkB;
        arrProduk[2] = produkC;
        
        for(int i = 0; i < arrProduk.length; i++){
            System.out.println(String.format("---Informasi Produk %s---", i+1));
            pM.tampilkanInformasiProduk(arrProduk[i]);
            System.out.println("\n");
        }
        
    }
}
