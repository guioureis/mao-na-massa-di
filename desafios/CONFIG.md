## Passo-a-passo para configurar seu acesso aos servidores:

1. Realize o *Fork* deste repositório para sua conta;
2. Edite o arquivo [users.yml](https://github.com/guitoper/mao-na-massa-di/blob/master/users.yml) e adcione as seguintes informações:
  * **username:** usuário que será utilizado para o login ssh. Recomento utilizar o mesmo usuário do seu Mac para não precisar informar o mesmo no comando de conexão ssh;
  * **pubkey:** chave ssh pública que permitirá seu acesso ao servidor. Utilize [este guia](https://git-scm.com/book/pt-br/v1/Git-no-Servidor-Gerando-Sua-Chave-P%C3%BAblica-SSH) caso não tenha uma chave ou não saiba o que é uma chave ssh;
  * **port:** porta que será utilizada para acessar sua aplicação. As portas disponíveis vão de 50000 até 60000. **Cuidado para não escolher uma porta que já está sendo utilizada**.
3. Faça um Pull Request da sua alteração e comunique no grupo do Slack. Por enquanto este processo está manual e alguém irá liberar seu acesso em nossos servidores.
4. Utilize o comando `ssh SEU_USUARIO@ec2-18-221-178-11.us-east-2.compute.amazonaws.com` para testar se seu acesso está liberado.
