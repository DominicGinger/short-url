FROM scratch

COPY main main 

EXPOSE 3003

CMD ["./main"]
