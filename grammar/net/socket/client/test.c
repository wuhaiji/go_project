#include<stdio.h>
#include<unistd.h>
#include<fcntl.h>
#include<string.h>
int main()
{
    int fd = open("./test.txt",O_RDWR|O_CREAT|O_APPEND,0777);
    if (fd < 0)
    {
        perror("error\n");
        return -1;
    }
    const char *ptr = "this is the input of text!";
    int ret = write(fd,ptr,strlen(ptr));
    if (ret < 0)
    {
        perror("error\n");
        return -1;
    }
    char buf[1024] = {0};
    lseek(fd,0,SEEK_SET);
    ret = read(fd,buf,1023);
    if (ret < 0)
    {
        perror("error!\n");
        return -1;
    }
    printf("%d-%s\n",ret,buf);
    close(fd);
    return 0;
}
