package main
import "fmt"

func main() {
    
  var solde int
  var retrait int
  
  fmt.Println("Donner le solde :")
  fmt.Scan(&solde)
  
  fmt.Println("Donner le retrait :")
  fmt.Scan(&retrait)
  
  if retrait <=0 {
      fmt.Println("Montant invalide")
      
  } else if retrait > solde {
      fmt.Println("Solde insuffisant")
    
  } else {
      fmt.Println("Retrait autorisé")
      fmt.Println("Nouveau solde :", solde-retrait)
  }
}
