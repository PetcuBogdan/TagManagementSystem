from dataclasses import field
from rest_framework import serializers
from account.models import User

class UserSerializer(serializers.ModelSerializer):
    password = serializers.CharField(write_only=True)
    confirm_password = serializers.CharField(write_only=True)

    class Meta:
        model = User
        fields = ["email", "name", "password", "confirm_password"]

    def validate(self, attrs):
        password = attrs.get('password')
        confirm_password = attrs.get('confirm_password')

        if password != confirm_password:
            raise serializers.ValidationError("Password and Confirm_Password doesn't match.")
        if len(password) < 8:
            raise serializers.ValidationError("Your password should have a minimum length of 8 characters")
        if not has_symbol(password):
            raise serializers.ValidationError("Your paassword should include at least one special character (e.g., !, @, #, $, %, ^, &, *)")    
        if not contains_number(password):       
            raise serializers.ValidationError("Your paassword should include at least one numerical digit (0-9)")
        if not contains_uppercase(password):
            raise serializers.ValidationError("Your password should include at least one uppercase letter (A-Z)")
        
        return attrs

    
    def validate_email(self, value):
        if User.objects.filter(email=value).exists():
          raise serializers.ValidationError('user with this Email already exists.')
        return value

    
    def create(self, validated_data):
        user = User.objects.create_user(
            email=validated_data['email'],
            name=validated_data['name'],
            password=validated_data['password'],
        )
        user.is_active = False
        user.save()
        return user
    
    def update(self, instance, validated_data):
        instance.name = validated_data.get('name', instance.name)
        instance.save()
        return instance

def has_symbol(s):
    symbols = "!@#$%^&*()_+{}[]|\;'\":<>,.?/~`"
    
    for symbol in symbols:
        if symbol in s:
            return True
    
    return False

def contains_number(s):
    for char in s:
        if char.isdigit():
            return True
    
    return False

def contains_uppercase(s):
    for char in s:
        if char.isupper():
            return True
    
    return False