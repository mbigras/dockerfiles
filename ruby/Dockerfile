FROM ruby:2.7
# throw errors if Gemfile has been modified since Gemfile.lock
RUN bundle config --global frozen 1
WORKDIR /app
COPY Gemfile Gemfile.lock /app/
RUN bundle install
COPY app.rb config.ru /app/
ENV PORT=3000
CMD bundle exec rackup -o 0.0.0.0 -p $PORT
