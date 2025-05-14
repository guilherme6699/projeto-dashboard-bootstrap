package main 

import(

	"fmt"
	"strings"
	   "os"
    "runtime"
	"os/exec"
)

func main(){

	nomeu := "gui"
	senhau := "46092560"

	fmt.Println("\n          MENU V1   \n")

	fmt.Print("Digite o Usuario:")
	var Usuario string
	fmt.Scanln(&Usuario)


	fmt.Println("Digite a senha:")
	var senha string
	fmt.Scanln(&senha)

	UsuarioCorreto := strings.TrimSpace(Usuario) == nomeu
	SenhaCorreta := strings.TrimSpace(senha) == senhau
	
	if UsuarioCorreto && SenhaCorreta {
		fmt.Println("\n ✅ Login Realizado com Sucesso!\n")
		fmt.Println("-------SELECIONE A OPÇÃO DESEJADA-------")
		fmt.Println("1 - Reiniciar o Serviço do xamp")
		fmt.Println("2 - ")

		var opcao int 
		_, err := fmt.Scanln(&opcao)
		if err != nil {
			fmt.Println("Erro, Digite um numero valido")
			return			
		}

		switch opcao{
	case  1:
		 xamp()
	case 2:
		fmt.Println("Saindo..")
		return
	default:
		fmt.Println("opção Invalida")
	  }
	} else{
		fmt.Println("\n❌ Login Falhou! Verifique o usuário ou senha.\n")
	}
}

		func xamp(){
			fmt.Println("🔄 Reiniciando serviços do XAMPP...")

	// Verifica o sistema operacional
	switch runtime.GOOS {
	case "linux":
		reiniciarXAMPPLinux()
	case "windows":
		reiniciarXAMPPWindows()
	default:
		fmt.Println("Sistema operacional não suportado.")
		os.Exit(1)
	}

	fmt.Println("✅ Serviços reiniciados com sucesso!")

		}
	func reiniciarXAMPPLinux() {
	cmds := []string{
		"sudo /opt/lampp/lampp stop",    // Para os serviços
		"sleep 2",                       // Espera 2 segundos
		"sudo /opt/lampp/lampp start",   // Reinicia
	}
	for _, cmd := range cmds{
		fmt.Println("Executando:", cmd)
		err := exec.Command("bash", "-c", cmd).Run()
		if err != nil {
			fmt.Printf("Erro ao executar o comando: '%s': %v\n",cmd, err)
			return
		}
	}
}
func reiniciarXAMPPWindows() {
	fmt.Println(" ⚠️comandos para Windows ainda não implementados")
		}
	
	



    
