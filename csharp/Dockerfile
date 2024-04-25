# Alpine Linux tabanlı .NET Core image
FROM mcr.microsoft.com/dotnet/sdk:8.0-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy csproj and restore as distinct layers
COPY *.csproj ./
RUN dotnet restore

# Copy the remaining files and build the application
COPY . ./
RUN dotnet publish -c Release -o out

# Build the runtime image
FROM mcr.microsoft.com/dotnet/runtime:8.0-alpine
WORKDIR /app
COPY --from=build /app/out .
ENTRYPOINT ["dotnet", "dotnet.dll"]
