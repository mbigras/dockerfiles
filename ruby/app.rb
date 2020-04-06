require 'sinatra'

get '/*' do |path|
  "URL Path = #{path}\n"
end
