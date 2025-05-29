package main 

import(

	"fmt"
	"strings"
	   "os"
    "runtime"
	"os/exec"
	"time"
)

func main(){

	nomeu := "gui"
	senhau := "46092560"
	fmt.Println("\n          MENU V1  ")
	fmt.Println("feito por:Guilherme/Cauan")

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
		fmt.Println("2 - Verificar o status dos Serviços")
		fmt.Println("3 - Ferramentas de Manutenção")
		fmt.Println("4 - Sair")
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

		fmt.Println("Verificando o status dos Serviços...")
		time.Sleep(2 * time.Second)
		fmt.Println("Serviços OK!✅")

	case 3:
		tool()
	case 4:
		fmt.Println("Saindo..")
		return


		

	default:
		fmt.Println("opção Invalida")
	  }
	} else{
		fmt.Println("\n❌ Login Falhou! Verifique o usuário ou senha.\n")
	}
}


func tool(){
	fmt.Println("------Ferramentas de Manutenção:------")
	fmt.Println("1 - vscode")
	fmt.Println("2 - xampp")
	fmt.Println("3 - Terminal")
	var opc int
	_, err := fmt.Scanln(&opc)
		if err != nil {
			fmt.Println("Erro, Digite um numero valido")
			return			
		}


	switch opc{

	case 1:
		code()

	case 2:
		time.Sleep(1 * time.Second)
		fmt.Println(" Em breve..")

	case 3:
	
	    term()
		default:
		fmt.Println("opção Invalida")
	}
}


func code(){
var cmd *exec.Cmd

switch runtime.GOOS {
	 case "windows":
		cmd = exec.Command("code", ".")
	 case "linux", "darwin":
		cmd = exec.Command("code",".")
	 default:
		fmt.Println("❌ Sistema operacional não suportado")
		os.Exit(1)
	}	
		fmt.Printf("🔄 Abrindo VS Code no diretori atual\n")
		err := cmd.Start()
		if err != nil {
			fmt.Printf("❌ Falha ao abrir VS Code: %v\n", err)
			fmt.Println("⚠️ Certifique-se que o VS Code está instalado e no PATH")
			os.Exit(1)
		}
		fmt.Println("✅ VS code aberto com sucesso")
		time.Sleep(2 * time.Second)
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

		func term() {

			var cmd *exec.Cmd

			switch runtime.GOOS{
			case "linux":
				cmd = exec.Command("gnome-terminal")
			case "windows":
				cmd = exec.Command("cmd", "/C", "start")
			default:
				fmt.Println("❌ Sistema Nâo Suportado")
				return 
			}

			err := cmd.Start()
			if err != nil {
				fmt.Printf("❌ Erro Ao Abrir o Terminal: %v")
			}else {
				fmt.Println("✅ Terminal Aberto com Sucesso")
			}
		}
	
	



    
